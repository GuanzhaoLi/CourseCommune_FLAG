import React from 'react';
import {List} from 'antd'
import axios from 'axios';
import Details from './QuestionDetails'


export default class CardList extends React.Component{

    state = {
        listData:[]
    }

    saveAnswers = (newData) => {
        this.setState({
            // eslint-disable-next-line no-undef
            listData:listData.map((item)=>{
                if(item.questionId===newData.questionId){
                    item.answer = newData.answer
                }
                //id不同就直接返回原来的值不修改，无论修改与否的的值都在这里返回
                return item
            })
        })
    }

    componentDidMount() {
        this.getQuestions();
    }

    getQuestions = () =>{
        //get data here
        const url = '';
        axios.get(url)
            .then(response =>{
                this.setState({
                listData: response.data.questions
            })
        }).catch((err)=>{
            console.log("ERROR",err.message);
            alert("Fail to get Data!");
        })
        // this.setState({
        //     listData:listData
        // })
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
                    key={item.questionId}
                    actions={[
                             <p>{item.subject}</p>,
                             <p>作业</p>,
                        // eslint-disable-next-line no-undef
                             <Details saveAnswers={this.saveAnswers} listData = {listData} data={item}></Details>
                    ]}
                >
                    <List.Item.Meta
                        title={<a className="titleOfList">待选问题{item.questionId} : {item.description}</a>}
                    />
                    <div className="contentOfQuestions">{item.content}</div>
                    <div style={{position: 'relative',fontSize:20, top: 5}} className="contentOfQuestions">回答： {item.answer}</div>

                </List.Item>
            )}
        />
    }
}
