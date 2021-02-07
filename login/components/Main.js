import React, { useState } from "react";
import { Route, Switch, Redirect } from "react-router";

import Login from "./Login";
import Register from "./Register";

function Main(props) {
    const { isLoggedIn, handleLoggedIn } = props;

    const showLogin = () => {
        return isLoggedIn ? (
            <Redirect to="/home" />
        ) : (
            <Login handleLoggedIn={handleLoggedIn} />
        );
    };


    return (
        <div className="main">
            <Switch>
                <Route path="/" exact render={showLogin} />
                <Route path="/login" render={showLogin} />
                <Route path="/register" component={Register} />
            </Switch>
        </div>
    );
}

export default Main;