import React from 'react';
import moment from 'moment';
import { ChatEntityType } from '../store/store';
import { UsernameWithContext } from '../shared/usernameWithContext';

const ServerMsgColor = '#237804';
const ServerErrColor = '#a8071a';

type EntityProps = {
  entityType: ChatEntityType;
  sender: string;
  senderId?: string;
  message: string;
  timestamp?: number;
  anonymous?: boolean;
};

export const ChatEntity = (props: EntityProps) => {
  let ts = '';
  if (props.timestamp) {
    ts = moment(props.timestamp).format('MMM Do - LT');
  }
  let el;
  switch (props.entityType) {
    case ChatEntityType.ServerMsg:
      el = (
        <div>
          <span style={{ color: ServerMsgColor }}>{props.message}</span>
        </div>
      );
      break;
    case ChatEntityType.ErrorMsg:
      el = (
        <div>
          <span style={{ color: ServerErrColor }}>{props.message}</span>
        </div>
      );
      break;
    case ChatEntityType.UserChat:
      el = (
        <div className="chat-entity">
          <p className="timestamp">{ts}</p>
          <p className="message-body">
            <span className="sender">
              {props.senderId && !props.anonymous ? (
                <UsernameWithContext
                  username={props.sender}
                  userID={props.senderId}
                />
              ) : (
                props.sender
              )}
            </span>
            <span className="message">{props.message}</span>
          </p>
        </div>
      );
      break;
    default:
      el = null;
  }
  return el;
};
