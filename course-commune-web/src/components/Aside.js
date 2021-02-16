import React from 'react';
import { Layout, Menu } from 'antd';
import { UserOutlined, LaptopOutlined, NotificationOutlined } from '@ant-design/icons';

const { SubMenu } = Menu;
const {Sider} = Layout;
function Aside(props) {
    return (
        <Sider width={300}
               style={{ background: '#d1dae2' }}
               className="site-layout-background">
            <Menu
                mode="inline"
                defaultSelectedKeys={['1']}
                defaultOpenKeys={['sub1']}
                style={{ height: '100%', borderRight: 0 }}
            >
                <SubMenu key="sub1" icon={<UserOutlined />} title="账户管理">
                    <Menu.Item key="1">用户中心</Menu.Item>
                    <Menu.Item key="2">个人历史</Menu.Item>
                </SubMenu>
                <SubMenu key="sub2" icon={<LaptopOutlined />} title="课程管理">
                    <Menu.Item key="5">账户信息</Menu.Item>
                    <Menu.Item key="3">查看提问</Menu.Item>
                    <Menu.Item key="4">查看视频预约</Menu.Item>
                    <Menu.Item key="6">查看已预约课程</Menu.Item>
                    <Menu.Item key="7">查看已接受问题</Menu.Item>
                </SubMenu>
                <SubMenu key="sub3" icon={<NotificationOutlined />} title="通知">
                    <Menu.Item key="9">消息</Menu.Item>
                    <Menu.Item key="10">通知</Menu.Item>
                </SubMenu>
            </Menu>
        </Sider>
    );
}

export default Aside;