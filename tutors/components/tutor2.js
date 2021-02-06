import React from 'react';
import {List} from 'antd'
import axios from 'axios';
import Details from './QuestionDetails'


// test data
const listData = [
    {
        href: 'https://ant.design',
        questionId: `0`,
        subject: '化学',
        studentId :'10086',
        description: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        content: '问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述',
        answer: '',
    },
    {
        href: 'https://ant.design',
        questionId: `1`,
        studentId: '10010',
        subject: '数学',
        description: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        content: '问题描述2',
        answer: '初始回答',
    }
];




export default class CardList extends React.Component{

    state = {
        listData:[]
    }

    // saveAnswers = (data) => {
    //     this.setState({
    //         listData : listData.map(data.questionId,data)
    //     })
    // }

    componentDidMount() {
        this.getQuestions();
    }

    getQuestions = () =>{
        //get data here
        // const url = BASE_URL;
        // axios.get(url).
        //     then(response =>{
        //         this.setState({
        //         listData: response.data.questions
        //     })
        // }).catch((err)=>{
        //     console.log("ERROR",err.message);
        //     alert("Fail to get Data!");
        // })
        this.setState({
            listData:listData
        })
    }

    render(){

        return <List
            className="whole_list"
            itemLayout="vertical"
            size="large"
            pagination={{
                onChange: page => {
                    console.log(page);
                },
                pageSize: 3,
            }}
            dataSource={this.state.listData}
            renderItem={item => (
                <List.Item className="listItems"
                    key={item.id}
                    actions={[
                             <p>{item.subject}</p>,
                             <p>作业</p>,
                             <Details listData = {listData} data={item}></Details>
                    ]}
                >
                    <List.Item.Meta
                        title={<a className="titleOfList" href={item.href}>待选问题{item.questionId}</a>}
                    />
                    <div className="contentOfQuestions">{item.content}</div>
                    <div className="contentOfQuestions">{item.answer}</div>

                </List.Item>
            )}
        />
    }
}
