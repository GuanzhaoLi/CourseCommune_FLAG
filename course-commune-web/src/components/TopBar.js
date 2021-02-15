import React from 'react';
import logo from "../assets/images/logo.svg"
import { Layout, Menu } from 'antd';
import Icon, { LogoutOutlined } from '@ant-design/icons';
import {Link} from "react-router-dom"
const {Header} = Layout;


function TopBar(props) {
    const { isLoggedIn, handleLogout } = props;

    return (
        <Header className="header">
            {/*<div className="logo" />*/}
            <img src={logo} className="app-logo" alt="logo" />
            <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['2']}>
                <Menu.Item key="1">
                    <Link to="/home">
                        <Icon type="fire" />
                        <span>首页</span>
                    </Link>
                </Menu.Item>
                <Menu.Item key="2">发现</Menu.Item>
                <Menu.Item key="3">提问</Menu.Item>
                <Menu.Item key="4">课程</Menu.Item>
                {isLoggedIn? null : (
                    <Menu.Item key="5" style={{float: 'right'}}>
                        <Link to="/register">
                            <Icon type="fire"/>
                            <span>注册</span>
                        </Link>
                    </Menu.Item>)}
                {isLoggedIn? null : (
                    <Menu.Item key="6" style={{float: 'right'}}>
                        <Link to="/login">
                        <Icon type="fire" />
                        <span>登录</span>
                        </Link>
                    </Menu.Item>) }
                {isLoggedIn ? (
                    <Menu.Item key="7" style={{float: 'right'}}>
                        <span style={{marginRight:"10px"}}>欢迎回来</span>
                        <LogoutOutlined className="logout" onClick={handleLogout} />
                    </Menu.Item>
                ) : null}
            </Menu>

        </Header>
    );
}

export default TopBar;