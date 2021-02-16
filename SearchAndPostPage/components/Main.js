import React from 'react';
import { Layout} from 'antd';
import { Input } from "antd";
import ListData from "./ListData";
const {Content} = Layout;
const { Search } = Input;


function Main(props) {



    return (
        <Content
            className="site-layout-background">
            <Search
                placeholder="input search text"
                enterButton="Search"
                size="large"
            />
            <ListData/>
        </Content>
    );
}

export default Main;