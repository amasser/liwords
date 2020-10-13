import React from 'react';

export const useDrawing = (isEnabled: boolean) => {
  // Drawing functionalities.
  // Right-drag = draw.
  // RightClick several times = clear drawing.
  // Shift+RightClick = clear drawing.
  // Shift+RightClick (when no drawing) = context menu.

  const boardEltRef = React.useRef<HTMLElement>();
  const [boardSize, setBoardSize] = React.useState({
    left: 0,
    top: 0,
    width: 1,
    height: 1,
  });
  const resizeFunc = React.useCallback(() => {
    const boardElt = boardEltRef.current;
    if (boardElt) {
      const { left, top, width, height } = boardElt.getBoundingClientRect();
      setBoardSize({
        left,
        top,
        width: Math.max(1, width),
        height: Math.max(1, height),
      });
    }
  }, []);
  const boardRef = React.useCallback(
    (elt) => {
      boardEltRef.current = elt;
      resizeFunc();
    },
    [resizeFunc]
  );
  React.useEffect(() => {
    if (!isEnabled) return;
    window.addEventListener('resize', resizeFunc);
    return () => window.removeEventListener('resize', resizeFunc);
  }, [isEnabled, resizeFunc]);
  const getXY = React.useCallback(
    (evt: React.MouseEvent): { x: number; y: number } => {
      const x = Math.max(
        0,
        Math.min(1, (evt.clientX - boardSize.left) / boardSize.width)
      );
      const y = Math.max(
        0,
        Math.min(1, (evt.clientY - boardSize.top) / boardSize.height)
      );
      return { x, y };
    },
    [boardSize]
  );
  const scaledXYStr = React.useCallback(
    ({ x, y }: { x: number; y: number }) =>
      `${x * boardSize.width},${y * boardSize.height}`,
    [boardSize.width, boardSize.height]
  );

  const boardResizedSinceLastPaintRef = React.useRef(true);
  const penRef = React.useRef<string>();
  const strokesRef = React.useRef<
    Array<{
      points: Array<{ x: number; y: number }>; // scaled to [0,1)
      path: string; // Mx,yLx,yLx,y... based on current boardSize
      pen: string; // "red"
      elt: React.ReactElement | undefined;
    }>
  >([]);
  const [currentDrawing, setCurrentDrawing] = React.useState();
  const plannedRepaintRef = React.useRef<number>();

  // For hopefully-unique id generation.
  const imagePrefixRef = React.useRef(`img${Date.now() + Math.random()}`);

  const repaintNow = React.useCallback(() => {
    if (plannedRepaintRef.current != null) {
      cancelAnimationFrame(plannedRepaintRef.current);
      plannedRepaintRef.current = undefined;
    }

    if (boardResizedSinceLastPaintRef.current) {
      // Rescale everything.
      for (const stroke of strokesRef.current) {
        let path = `M${scaledXYStr(stroke.points[0])}`;
        for (let i = 1; i < stroke.points.length; ++i) {
          path += `L${scaledXYStr(stroke.points[i])}`;
        }
        stroke.path = path;
        stroke.elt = undefined;
      }
      boardResizedSinceLastPaintRef.current = false;
    }

    for (let i = 0; i < strokesRef.current.length; ++i) {
      const stroke = strokesRef.current[i];
      if (!stroke.elt) {
        let path = stroke.path;
        if (stroke.points.length === 1) {
          // Draw a diamond to represent a single point.
          path += 'm-1,0l1,1l1,-1l-1,-1l-1,1l1,1';
        }
        stroke.elt =
          stroke.pen === 'erase' ? (
            <path key={i} d={path} fill="none" strokeWidth={5} stroke="black" />
          ) : (
            <path
              key={i}
              d={path}
              fill="none"
              strokeWidth={5}
              stroke={stroke.pen}
            />
          );
      }
    }

    let toDraw: Array<React.ReactElement> = [];
    let toErase: Array<React.ReactElement> = [];
    const eraseMasks: Array<React.ReactElement> = [];
    let numMasks = 0;
    for (let i = 0; i < strokesRef.current.length; ) {
      for (
        ;
        i < strokesRef.current.length && strokesRef.current[i].pen !== 'erase';
        ++i
      ) {
        toDraw.push(strokesRef.current[i].elt!);
      }
      for (
        ;
        i < strokesRef.current.length && strokesRef.current[i].pen === 'erase';
        ++i
      ) {
        toErase.push(strokesRef.current[i].elt!);
      }
      if (toErase.length > 0) {
        if (toDraw.length > 0) {
          // Otherwise nothing worth erasing.
          const maskId = `${imagePrefixRef.current}.${++numMasks}`;
          eraseMasks.push(
            <mask key={i - 1} id={maskId}>
              <rect
                width={boardSize.width}
                height={boardSize.height}
                fill="white"
              />
              {toErase}
            </mask>
          );
          toDraw = [
            <g key={i - 1} mask={`url(#${maskId})`}>
              {toDraw}
            </g>,
          ];
        }
        toErase = [];
      } else if (toDraw.length > 0) {
        toDraw = [<g key={i - 1}>{toDraw}</g>];
      }
    }

    const ret = (
      <React.Fragment>
        {eraseMasks}
        {toDraw}
      </React.Fragment>
    );

    setCurrentDrawing(ret);
  }, [scaledXYStr, boardSize.width, boardSize.height]);

  const scheduleRepaint = React.useCallback(() => {
    if (plannedRepaintRef.current != null) {
      cancelAnimationFrame(plannedRepaintRef.current);
    }
    plannedRepaintRef.current = requestAnimationFrame(repaintNow);
  }, [repaintNow]);

  const handleContextMenu = React.useCallback(
    (evt: React.MouseEvent) => {
      if (!evt.shiftKey) {
        // Draw when not holding shift.
        evt.preventDefault();
      } else if (strokesRef.current.length > 0) {
        // Shift+RightClick clears drawing.
        evt.preventDefault();
        penRef.current = undefined;
        strokesRef.current = [];
        scheduleRepaint();
      } else {
        // Shift+RightClick accesses context menu if no drawing.
      }
    },
    [scheduleRepaint]
  );

  const handleMouseDown = React.useCallback(
    (evt: React.MouseEvent) => {
      if (evt.button === 2 && !evt.shiftKey) {
        const newXY = getXY(evt);
        penRef.current = strokesRef.current.length % 3 === 2 ? 'erase' : 'red';
        strokesRef.current.push({
          points: [newXY],
          path: `M${scaledXYStr(newXY)}`,
          pen: penRef.current,
          elt: undefined,
        });
        scheduleRepaint();
      }
    },
    [getXY, scaledXYStr, scheduleRepaint]
  );

  const handleMouseUp = React.useCallback(
    (evt: React.MouseEvent) => {
      if (!penRef.current) return;
      // Right-click this many times to clear drawing.
      const howMany = 3;
      if (strokesRef.current.length >= howMany) {
        const lastPoint =
          strokesRef.current[strokesRef.current.length - 1].points[0];
        let i = 0;
        for (; i < howMany; ++i) {
          const ithLastPoints =
            strokesRef.current[strokesRef.current.length - (i + 1)].points;
          if (
            !(
              ithLastPoints.length < 2 &&
              ithLastPoints[0].x === lastPoint.x &&
              ithLastPoints[0].y === lastPoint.y
            )
          )
            break;
        }
        if (i === howMany) strokesRef.current = [];
      }
      penRef.current = undefined;
      scheduleRepaint();
    },
    [scheduleRepaint]
  );

  const handleMouseMove = React.useCallback(
    (evt: React.MouseEvent) => {
      if (!penRef.current) return;
      const newXY = getXY(evt);
      const lastStroke = strokesRef.current[strokesRef.current.length - 1];
      const lastPoints = lastStroke.points;
      const lastPoint = lastPoints[lastPoints.length - 1];
      if (lastPoint.x === newXY.x && lastPoint.y === newXY.y) return;
      lastPoints.push(newXY);
      lastStroke.path += `L${scaledXYStr(newXY)}`;
      lastStroke.elt = undefined; // will be recomputed later
      scheduleRepaint();
    },
    [getXY, scaledXYStr, scheduleRepaint]
  );

  const handlePointerDown = React.useCallback((evt: React.PointerEvent) => {
    (evt.target as Element).setPointerCapture(evt.pointerId);
  }, []);
  const handlePointerUp = React.useCallback(
    (evt: React.PointerEvent) => {
      (evt.target as Element).releasePointerCapture(evt.pointerId);
      handleMouseUp(evt);
    },
    [handleMouseUp]
  );

  React.useEffect(() => {
    // Board size changed, invalidate path.
    if (strokesRef.current.length > 0) {
      boardResizedSinceLastPaintRef.current = true;
      scheduleRepaint();
    }
  }, [scaledXYStr, scheduleRepaint]);

  const outerDivProps = React.useMemo(
    () =>
      isEnabled
        ? {
            ref: boardRef,
            onContextMenu: handleContextMenu,
            onMouseDown: handleMouseDown,
            onMouseUp: handleMouseUp,
            onMouseMove: handleMouseMove,
            onPointerDown: handlePointerDown,
            onPointerUp: handlePointerUp,
          }
        : {},
    [
      isEnabled,
      boardRef,
      handleContextMenu,
      handleMouseDown,
      handleMouseUp,
      handleMouseMove,
      handlePointerDown,
      handlePointerUp,
    ]
  );
  const svgProps: React.SVGProps<SVGSVGElement> = React.useMemo(
    () =>
      isEnabled
        ? {
            viewBox: `0 0 ${boardSize.width} ${boardSize.height}`,
            style: {
              position: 'absolute',
              left: 0,
              top: 0,
              width: boardSize.width,
              height: boardSize.height,
              pointerEvents: 'none',
            },
          }
        : {},
    [isEnabled, boardSize.width, boardSize.height]
  );
  const svgDrawing = React.useMemo(
    () =>
      isEnabled && currentDrawing ? (
        <svg {...svgProps}>{currentDrawing}</svg>
      ) : null,
    [isEnabled, svgProps, currentDrawing]
  );
  const ret = React.useMemo(() => ({ outerDivProps, svgDrawing }), [
    outerDivProps,
    svgDrawing,
  ]);
  return ret;
};
