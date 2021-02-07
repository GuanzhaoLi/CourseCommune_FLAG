import React from 'react';
import 'antd/dist/antd.css';
import { Modal, Button } from 'antd';
import axios from 'axios';

export default class Details extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            isModalVisible : false,
            data: this.props.data,
            textareaValue: this.props.data.answer,
        }
    }

    //Clicking cancel or exit won't save the answer
    //ShowModal will call the question Details with answer input
    showModal = () => {
        this.setState({
            isModalVisible : true
        })
    };

    // Submit the answer from tutor here
    // Answer will be submitted with clicking ok
    Submit = () => {
        const opt = {
            method: "POST",
            url: `/answer`,
            data: {
                answer: this.state.textareaValue
            },
        }
        axios(opt)
            .then((res) => {
                if (res.status === 200) {
                    console.log(res.data);
                }
            })
            .catch((err) => {
                console.log("答案上传错误 ", err.message);
                alert("Error happens, Please update the answer again!");
            });
        console.log(this.state.textareaValue);

        const data = {...this.state.data}
        data.answer = this.state.textareaValue


        this.props.saveAnswers(data);

        // Close the window
        this.setState({
            data: data,
            isModalVisible : false
        })
    };

    handleCancel = () => {
        this.setState({
            isModalVisible : false
        })
    };

    // automatically save the change of textarea
    onInputChange = (e) => {
        this.setState({
            textareaValue:e.target.value
        })
    }

    render(){
        return (
            <>
                <Button type="primary" onClick={this.showModal}>
                    回答
                </Button>
                <Modal width={1000} title="问题详情" visible={this.state.isModalVisible} onOk={this.Submit} onCancel={this.handleCancel}>
                    <p>问题：{this.props.data.questionId}.{this.props.data.description}</p>
                    <p>问题人：{this.props.data.studentId}</p>
                    <p>内容：{this.props.data.content}</p>
                    <p>回答：</p>
                    <textarea className="answerOfQuestion" onChange={this.onInputChange}></textarea>
                </Modal>
            </>
        );
    }
}
