import React, { useCallback, useEffect, useState } from 'react';
import { Route, Switch } from 'react-router-dom';
import './App.scss';
import axios from 'axios';
import 'antd/dist/antd.css';

import { Table } from './gameroom/table';
import { Lobby } from './lobby/lobby';
import {
  useExcludedPlayersStoreContext,
  useRedirGameStoreContext,
} from './store/store';

import { LiwordsSocket } from './socket/socket';
import { About } from './about/about';
import { Login } from './lobby/login';
import { Register } from './lobby/register';
import { UserProfile } from './profile/profile';
import { PasswordChange } from './lobby/password_change';
import { PasswordReset } from './lobby/password_reset';
import { NewPassword } from './lobby/new_password';
import { toAPIUrl } from './api/api';

type Blocks = {
  user_ids: Array<string>;
};

const App = React.memo(() => {
  const { setExcludedPlayers } = useExcludedPlayersStoreContext();
  const { redirGame, setRedirGame } = useRedirGameStoreContext();
  const [shouldDisconnect, setShouldDisconnect] = useState(false);

  const [liwordsSocketValues, setLiwordsSocketValues] = useState({
    sendMessage: (msg: Uint8Array) => {},
    justDisconnected: false,
  });
  const { sendMessage } = liwordsSocketValues;

  if (redirGame !== '') {
    setRedirGame('');
    window.location.replace(`/game/${redirGame}`);
  }

  const disconnectSocket = useCallback(() => {
    setShouldDisconnect(true);
    setTimeout(() => {
      // reconnect after 5 seconds.
      setShouldDisconnect(false);
    }, 5000);
  }, []);

  useEffect(() => {
    axios
      .post<Blocks>(
        toAPIUrl('user_service.SocializeService', 'GetFullBlocks'),
        {},
        { withCredentials: true }
      )
      .then((resp) => {
        setExcludedPlayers(new Set<string>(resp.data.user_ids));
      });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <div className="App">
      <LiwordsSocket
        disconnect={shouldDisconnect}
        setValues={setLiwordsSocketValues}
      />
      <Switch>
        <Route path="/" exact>
          <Lobby sendSocketMsg={sendMessage} DISCONNECT={disconnectSocket} />
        </Route>
        <Route path="/game/:gameID">
          {/* Table meaning a game table */}
          <Table sendSocketMsg={sendMessage} />
        </Route>
        <Route path="/about">
          <About />
        </Route>
        <Route path="/login">
          <Login />
        </Route>
        <Route path="/secretwoogles">
          <Register />
        </Route>
        <Route path="/password/change">
          <PasswordChange />
        </Route>
        <Route path="/password/reset">
          <PasswordReset />
        </Route>

        <Route path="/password/new">
          <NewPassword />
        </Route>

        <Route path="/profile/:username">
          <UserProfile />
        </Route>
      </Switch>
    </div>
  );
});

export default App;
