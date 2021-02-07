import React from 'react';
import 'antd/dist/antd.css';
import {Modal, Button} from 'antd';
import axios from 'axios';

export default class VideoDetails extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            isModalVisible : false,
            data: this.props.data,
        }
    }

    showModal = () => {
        this.setState({
            isModalVisible : true
        })
    };

    accept = () => {
        const opt = {
            method: "POST",
            url: `/video`,
            data: {
                data: this.state.data
            },
        }
        axios(opt)
            .then((res) => {
                if (res.status === 200) {
                    console.log(res.data);
                }
            })
            .catch((err) => {
                console.log("预约失败", err.message);
                alert("Error happens, Please accept again!");
            });

        const data = {...this.state.data}
        data.accepted = !data.accepted

        this.props.update(data);

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
        const item = this.props.data;
        return (
            <>
                <Button type="primary" onClick={this.showModal}>
                    {item.accepted==false? "预约":"取消预约"}
                </Button>
                <Modal width={1000} title="问题详情" visible={this.state.isModalVisible}
                       onOk={this.accept} onCancel={this.handleCancel}>
                    <p>视频任务{item.id}</p>
                    <p>日期：{item.startDate}-{item.endDate}</p>
                    <p>学生：{item.avatar}</p>
                    <p>学科: {item.subject}</p>
                    <p>描述: {item.content}</p>
                    <p>视频任务状态：{item.accepted==false? "未预约":"已预约"}</p>
                    <p style={{fontSize:20, textAlign:'center'}}>{item.accepted==false? "你确定接受此视频任务吗?":"你确定取消此预约吗?"}</p>
                </Modal>
            </>
        );
    }
}