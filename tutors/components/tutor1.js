import React, { Component } from 'react'
import Masonry from 'masonry-layout'  //实现瀑布流
import imagesloaded  from 'imagesloaded' //监听图片加载
import InfiniteScroll from 'react-infinite-scroller' //下拉加载
import axios from 'axios'

import {Button, Card, Modal} from 'antd'
import BASE_URL from '../constants'
import {EllipsisOutlined} from '@ant-design/icons'



// get data here
// test data
const arr = [
    {
        href: 'https://ant.design',
        id: `0`,
        startDate: "2.6 19:00",
        endDate: "2.6 20:00",
        subject: '学科',
        avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        content: '问题描述',
        acceptOrNot: false
    },
    {
        href: 'https://ant.design',
        id: `0`,
        subject: '学科',
        avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        content: '问题描述'
    },
    {
        href: 'https://ant.design',
        id: `0`,
        subject: '学科',
        avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        content: '问题描述'
    },
    {
        href: 'https://ant.design',
        id: `0`,
        subject: '学科',
        avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        content: '问题描述'
    },
    {
        href: 'https://ant.design',
        id: `0`,
        subject: '学科',
        avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        content: '问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述问题描述'
    },
    {
        href: 'https://ant.design',
        id: `0`,
        subject: '学科',
        avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        content: '问题描述'
    },
    {
        href: 'https://ant.design',
        id: `0`,
        subject: '学科',
        avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        content: '问题描述'
    },
    {
        href: 'https://ant.design',
        id: `0`,
        subject: '学科',
        avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        content: '问题描述'
    },
    {
        href: 'https://ant.design',
        id: `0`,
        subject: '学科',
        avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        content: '问题描述'
    },
    {
        href: 'https://ant.design',
        id: `0`,
        subject: '学科',
        avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        content: '问题描述'
    },
]

export default class Pinterest extends Component {
    state = {
        data: [],
        hasMore: true,
        isModalVisible:false
    }

    getVideoData = ()=> {
        // define the url here
        // const url = BASE_URL;
        // axios.get(url).
        // then(response =>{
        //     this.setState({
        //         data: response.data.videos
        //     })
        // }).catch((err)=>{
        //     console.log("ERROR",err.message);
        //     alert("Fail to get Data!");
        // })
        this.setState({
            data : arr,
        })
    }

    componentDidMount () {
        // get data here
        this.getVideoData();
        this.imagesOnload()
    }

    showModal = () => {
        this.setState({
            isModalVisible : true
        })
    };

    handleCancel = () => {
        this.setState({
            isModalVisible : false
        })
    };

    accept = () =>{
        // when tutor accept to get the video request, click ok
        //update the information here

        const opt = {
            method: "POST",
            url: `${BASE_URL}/accept`,
            data: {
                acceptOrNot: true
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


        alert("已接受预约");
        this.setState({
            isModalVisible : false
        })
    }

    //Lazy loading
    imagesOnload = () => {
        const elLoad = imagesloaded('.pages_hoc')
        elLoad.on('always', () => {
            // 调用瀑布流
            this.advanceWidth()
        })
    }

    //瀑布流
    advanceWidth = () => {

        var elem = document.querySelector('.pages_hoc');

        new Masonry( elem, {
            itemSelector: '.imgBox',
            columnWidth: '.imgBox',
            fitWidth: true,
            gutter: 20,
        });
    }


    render() {
        const {hasMore} = this.state
        let data = Array.from(this.state.data);
        return (
            <div className='pages_pinterest'>
                {/* 下拉加载 */}
                <InfiniteScroll
                    initialLoad={false}
                    hasMore={hasMore}
                    useWindow={false}
                >
                    <div className="pages_hoc">
                        {
                            data.map((item, index) => {
                                if(item.acceptOrNot==false) {
                                    return (
                                        <div key={index} className='imgBox'>
                                            <Card className="cards">
                                                <p>视频任务{index}</p>
                                                <p>日期：{item.startDate}-{item.endDate}</p>
                                                <p>学生：{item.avatar}</p>
                                                <p>学科: {item.subject}</p>
                                                <p>描述: {item.content}</p>
                                                <Button type="primary" onClick={this.showModal}>
                                                    预约
                                                </Button>
                                            </Card>
                                            <Modal width={1000} title="问题详情" visible={this.state.isModalVisible}
                                                   onOk={this.accept} onCancel={this.handleCancel}>
                                                <p>视频任务{index}</p>
                                                <p>日期：{item.startDate}-{item.endDate}</p>
                                                <p>学生：{item.avatar}</p>
                                                <p>学科: {item.subject}</p>
                                                <p>描述: {item.content}</p>
                                                <p>你确定接受此视频任务吗?</p>
                                            </Modal>
                                        </div>
                                    )
                                }
                            })
                        }
                    </div>
                </InfiniteScroll>
            </div>
        )
    }
}