import React from 'react';
import { Breadcrumb } from 'antd';
function Crumb(props) {
    return (
        <Breadcrumb style={{ margin: '16px 0' }}>
            <Breadcrumb.Item>搜索问题</Breadcrumb.Item>
            <Breadcrumb.Item>现有问题</Breadcrumb.Item>
            <Breadcrumb.Item>App</Breadcrumb.Item>
        </Breadcrumb>
    );
}

export default Crumb;