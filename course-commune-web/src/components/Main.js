import React, { useState } from "react";
import { Route, Switch, Redirect } from "react-router";

import Login from "./Login";
import Register from "./Register";
import Home from "./Home";
import { Layout} from 'antd';
import { Input } from "antd";

const {Content} = Layout;
const { Search } = Input;

function Main(props) {
    const { isLoggedIn, handleLoggedIn } = props;

    const showLogin = () => {
        return isLoggedIn ? (
            <Redirect to="/home" />
        ) : (
            <Login handleLoggedIn={handleLoggedIn} />
        );
    };
    const showHome = () => {
        return isLoggedIn ? <Home /> : <Redirect to="/login" />;

        // for test propose to show Home组件
        // return <Home />
    };

    return (
        <div>
            <Switch>
                <Route path="/" exact render={showLogin} />
                <Route path="/login" render={showLogin} />
                <Route path="/register" component={Register} />
                <Route path="/home" render={showHome} />
            </Switch>
        </div>
    );
}

export default Main;