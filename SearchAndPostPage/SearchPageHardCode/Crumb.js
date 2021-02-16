import React from 'react';
import { Breadcrumb } from 'antd';
function Crumb(props) {
    return (
        <Breadcrumb style={{ margin: '16px 0' }}>
            <Breadcrumb.Item>Search</Breadcrumb.Item>
            <Breadcrumb.Item>List</Breadcrumb.Item>
        </Breadcrumb>
    );
}

export default Crumb;