import React, { useState } from 'react';
import ReactDOM from 'react-dom';
import 'antd/dist/antd.css';
import axios from "axios";
import { BASE_URL } from "../constants";
import {
    Form,
    Input,
    Select,
    Row,
    Col,
    Checkbox,
    Button,
    message
} from 'antd';
import logo from "../assets/images/logo.svg";


const { Option } = Select;
const formItemLayout = {

    wrapperCol: {
        xs: { span: 8, offset: 8 },
        sm: { span: 8, offset: 8 }
    }
};
const tailFormItemLayout = {
    wrapperCol: {
        xs: {
            span: 8,
            offset: 10
        },
        sm: {
            span: 8,
            offset: 10
        }
    }
};

function Register (props) {
    const [form] = Form.useForm();

    const onFinish = (values) => {
        console.log('Received values of form: ', values);
        const { username, password } = values;
        const opt = {
            method: "POST",
            url: `${BASE_URL}/signup`,
            data: {
                username: username,
                password: password
            },
            headers: { "content-type": "application/json" }
        };

        axios(opt)
            .then((response) => {
                console.log(response);
                // case1: registered success
                if (response.status === 200) {
                    message.success("Registration succeed!");
                    props.history.push("/login");
                }
            })
            .catch((error) => {
                console.log("register failed: ", error.message);
                message.error("Registration failed!");
                // throw new Error('Signup Failed!')
            });
    };

    const prefixSelector = (
        <Form.Item name="prefix" noStyle>
            <Select
                style={{
                    width: 70,
                }}
            >
                <Option value="86">+86</Option>
                <Option value="87">+87</Option>
            </Select>
        </Form.Item>
    );

    return (
        <div className='register-block'>
            <Form
                {...formItemLayout}
                form={form}
                name="register-form"
                onFinish={onFinish}
                initialValues={{
                    prefix: '86',
                }}
                scrollToFirstError
            >
                {/*<img src={logo} className="register-logo"/>*/}
                <Form.Item
                    name="nickname"
                    rules={[
                        {
                            required: true,
                            message: '请输入昵称!',
                            whitespace: true,
                        },
                    ]}
                >
                    <Input placeholder="请输入昵称"/>
                </Form.Item>

                <Form.Item
                    name="email"
                    rules={[
                        {
                            type: 'email',
                            message: '邮箱格式不正确!',
                        },
                        {
                            required: true,
                            message: '请输入邮箱!',
                        },
                    ]}
                >
                    <Input placeholder="请输入邮箱"/>
                </Form.Item>

                <Form.Item
                    name="password"
                    rules={[
                        {
                            required: true,
                            message: '请输入密码!',
                        },
                    ]}
                    hasFeedback
                >
                    <Input.Password placeholder="请输入密码"/>
                </Form.Item>

                <Form.Item
                    name="confirm"
                    dependencies={['password']}
                    hasFeedback
                    rules={[
                        {
                            required: true,
                            message: '请确认密码!',
                        },
                        ({ getFieldValue }) => ({
                            validator(_, value) {
                                if (!value || getFieldValue('password') === value) {
                                    return Promise.resolve();
                                }
                                return Promise.reject('密码不一致!');
                            },
                        }),
                    ]}
                >
                    <Input.Password placeholder="请输入密码"/>
                </Form.Item>


                <Form.Item
                    name="phone"
                    rules={[
                        {
                            required: true,
                            message: 'Please input your phone number!',
                        },
                    ]}
                >
                    <Input
                        addonBefore={prefixSelector}
                        style={{
                            width: '100%',
                        }}
                        placeholder="请输入电话号码"
                    />
                </Form.Item>

                <Form.Item className="captcha">
                    <Row gutter={8}>
                        <Col span={12}>
                            <Form.Item
                                name="captcha"
                                noStyle
                                rules={[
                                    {
                                        required: true,
                                        message: 'Please input the captcha you got!',
                                    },
                                ]}
                            >
                                <Input placeholder="请输入验证码"/>
                            </Form.Item>
                        </Col>
                        <Col span={12}>
                            <Button>发送验证码</Button>
                        </Col>
                    </Row>
                </Form.Item>

                <Form.Item
                    name="agreement"
                    valuePropName="checked"
                    rules={[
                        {
                            validator: (_, value) =>
                                value ? Promise.resolve() : Promise.reject('Should accept agreement'),
                        },
                    ]}
                    {...tailFormItemLayout}
                >
                    <Checkbox>
                        I have read the <a href="">agreement</a>
                    </Checkbox>
                </Form.Item>

                <Form.Item {...tailFormItemLayout}>
                    <Button type="primary" htmlType="submit">
                        注册
                    </Button>
                </Form.Item>
            </Form>
        </div>

    );
};

export default Register;