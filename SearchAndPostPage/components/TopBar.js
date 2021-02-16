import React from 'react';
import logo from "../image/logo.svg"
import { Layout, Menu } from 'antd';
const {Header} = Layout;
function TopBar(props) {
    return (
        <Header className="header">
            {/*<div className="logo" />*/}
            <img src={logo} className="logo" alt="logo" />
            <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['2']}>
                <Menu.Item key="1">nav 1</Menu.Item>
                <Menu.Item key="2">nav 2</Menu.Item>
                <Menu.Item key="3">nav 3</Menu.Item>
            </Menu>
        </Header>
    );
}

export default TopBar;