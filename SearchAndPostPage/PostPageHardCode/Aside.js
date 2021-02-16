import React from 'react';
import { Layout, Menu } from 'antd';
import { UserOutlined, LaptopOutlined, SearchOutlined } from '@ant-design/icons';

const { SubMenu } = Menu;
const {Sider} = Layout;
function Aside(props) {
    return (
        <Sider width={200} className="site-layout-background">
            <Menu
                mode="inline"
                defaultSelectedKeys={['1']}
                defaultOpenKeys={['sub3']}
                style={{ height: '100%', borderRight: 0 }}
            >
                <SubMenu key="sub1" icon={<UserOutlined />} title="User Info">
                    <Menu.Item key="1">Account Info</Menu.Item>
                    <Menu.Item key="2">Switch Account</Menu.Item>
                    <Menu.Item key="3">Log out</Menu.Item>
                </SubMenu>
                <SubMenu key="sub2" icon={<LaptopOutlined />} title="Post">
                    <Menu.Item key="4">History</Menu.Item>
                    <Menu.Item key="5">Post Question</Menu.Item>
                </SubMenu>
                <SubMenu key="sub3" icon={<SearchOutlined />} title="Search">
                    <Menu.Item key="6">Search Question</Menu.Item>
                </SubMenu>
            </Menu>
        </Sider>
    );
}

export default Aside;