import React, { useEffect, useState } from 'react';
import { Route, Switch, useLocation, Redirect } from 'react-router-dom';
import './App.scss';
import 'antd/dist/antd.css';
import useWebSocket from 'react-use-websocket';

import { Table } from './gameroom/table';
import { Lobby } from './lobby/lobby';
import { useStoreContext } from './store/store';

import { getSocketURI } from './socket/socket';
import { decodeToMsg, encodeToSocketFmt } from './utils/protobuf';
import { onSocketMsg } from './store/socket_handlers';
import { About } from './about/about';
import { Login } from './lobby/login';
import { Register } from './lobby/register';
import { MessageType, JoinPath } from './gen/api/proto/realtime/realtime_pb';
import { useSocketToken } from './hooks/use_socket_token';
import { UserProfile } from './profile/profile';
import { PasswordChange } from './lobby/password_change';

const JoinSocketDelay = 1000;

const App = React.memo(() => {
  const store = useStoreContext();
  const socketUrl = getSocketURI();
  const [connectedToSocket, setConnectedToSocket] = useState(false);
  const { sendMessage } = useWebSocket(socketUrl, {
    onOpen: () => {
      console.log('connected to socket');
      setConnectedToSocket(true);
    },
    onClose: () => {
      console.log('disconnected from socket :(');
      setConnectedToSocket(false);
    },
    retryOnError: true,
    // Will attempt to reconnect on all close events, such as server shutting down
    shouldReconnect: (closeEvent) => true,
    onMessage: (event: MessageEvent) =>
      decodeToMsg(event.data, onSocketMsg(username, store)),
  });

  const { username, userID, loggedIn } = useSocketToken(
    sendMessage,
    connectedToSocket
  );
  const location = useLocation();

  useEffect(() => {
    console.log('location pathname change, now', location.pathname);
    const rr = new JoinPath();
    rr.setPath(location.pathname);
    console.log('Tryna register with path', location.pathname);
    setTimeout(() => {
      sendMessage(
        encodeToSocketFmt(MessageType.JOIN_PATH, rr.serializeBinary())
      );
    }, JoinSocketDelay);

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [location.pathname]);

  if (store.redirGame !== '') {
    store.setRedirGame('');
    return <Redirect push to={`/game/${store.redirGame}`} />;
  }

  return (
    <div className="App">
      <Switch>
        <Route path="/" exact>
          <Lobby
            username={username}
            userID={userID}
            sendSocketMsg={sendMessage}
            loggedIn={loggedIn}
            // connectedToSocket={connectedToSocket}
          />
        </Route>
        <Route path="/game/:gameID">
          {/* Table meaning a game table */}
          <Table
            sendSocketMsg={sendMessage}
            username={username}
            loggedIn={loggedIn}
            // can use some visual indicator to show the user if they disconnected
            // connectedToSocket={connectedToSocket}
          />
        </Route>
        <Route path="/about">
          <About myUsername={username} loggedIn={loggedIn} />
        </Route>
        <Route path="/login">
          <Login />
        </Route>
        <Route path="/secretwoogles">
          <Register />
        </Route>
        <Route path="/password/change">
          <PasswordChange username={username} loggedIn={loggedIn} />
        </Route>
        <Route path="/profile/:username">
          <UserProfile myUsername={username} loggedIn={loggedIn} />
        </Route>
      </Switch>
    </div>
  );
});

export default App;
