import React from 'react';
import { Layout} from 'antd';
import { Input } from "antd";
import ListData from "./ListData";
import PostForm from './PostForm';

const {Content} = Layout;
const { Search } = Input;


function Main(props) {



    return (
        <Content>
            <PostForm></PostForm>
        </Content>
    );
}

export default Main;