import React, {Component} from 'react';
import { Form, Input, Button, Checkbox, message } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import { Link } from "react-router-dom";
import axios from "axios";
import { BASE_URL } from "../constants";
import logo from "../assets/images/logo.svg"

function Login (props) {
    const { handleLoggedIn } = props;

    const onFinish = (values) => {
        console.log('Received values of form: ', values);
        const { username, password } = values;
        const opt = {
            method: "POST",
            url: `${BASE_URL}/signin`,
            data: {
                username: username,
                password: password
            },
            headers: { "Content-Type": "application/json" }
        };
        axios(opt)
            .then((res) => {
                if (res.status === 200) {
                    // console.log(res.data);
                    const { data } = res;
                    handleLoggedIn(data);
                    message.success("Login succeed! ");
                }
            })
            .catch((err) => {
                console.log("login failed: ", err.message);
                message.error("Login failed!");
            });
    };

    return (
        <div className="login-block">
            <Form
                name="normal_login"
                className="login-form"
                initialValues={{
                    remember: true,
                }}
                onFinish={onFinish}
            >
                <img src={logo} className="logo"/>
                <Form.Item
                    name="username"
                    rules={[
                        {
                            required: true,
                            message: 'Please input your Username!',
                        },
                    ]}
                >
                    <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="Username" />
                </Form.Item>
                <Form.Item
                    name="password"
                    rules={[
                        {
                            required: true,
                            message: 'Please input your Password!',
                        },
                    ]}
                >
                    <Input
                        prefix={<LockOutlined className="site-form-item-icon" />}
                        type="password"
                        placeholder="Password"
                    />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit" className="login-form-button">
                        登录
                    </Button>

                </Form.Item>
                <Form.Item>
                    <Form.Item name="remember" valuePropName="checked" noStyle>
                        <Checkbox>记住登录状态</Checkbox>
                    </Form.Item>
                    <a className="login-form-forgot" href="">
                        忘记密码？
                    </a>
                    <Link className="register-new-account" to="/register">注册新用户</Link>
                </Form.Item>
            </Form>
        </div>

    );
};



export default Login;



