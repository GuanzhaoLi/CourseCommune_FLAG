import React, { Component } from 'react'
import Masonry from 'masonry-layout'  //实现瀑布流
import imagesloaded  from 'imagesloaded' //监听图片加载
import InfiniteScroll from 'react-infinite-scroller' //下拉加载
import axios from 'axios'

import {Card} from 'antd'
import VideoDetails from './videoDetails'

export default class Pinterest extends Component {
    state = {
        data: [],
        hasMore: true,
    }

    getVideoData = ()=> {
        // define the url here
        const url = "";
        axios.get(url).
        then(response =>{
            this.setState({
                data: response.data.videos
            })
        }).catch((err)=>{
            console.log("ERROR",err.message);
            alert("Fail to get Data!");
        })
    }

    componentDidMount () {
        // get data here
        this.getVideoData();
        this.imagesOnload()
    }

    update = (newData) => {
        const data = this.state.data
        this.setState({
            data: data.map((item)=>{
                if(item.id===newData.id){
                    item.accepted = newData.accepted
                }
                return item
            })
        })
    }

    //Lazy loading
    imagesOnload = () => {
        const elLoad = imagesloaded('.pages_hoc')
        elLoad.on('always', () => {
            this.advanceWidth()
        })
    }

    //Pinterest Loading
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
                <InfiniteScroll
                    initialLoad={false}
                    hasMore={hasMore}
                    useWindow={false}
                >
                    <div className="pages_hoc">
                        {
                            data.map((item, index) => {
                                    return (
                                        <div key={index} className='imgBox'>
                                            <Card className="cards" actions={[<VideoDetails update={this.update} data={item}></VideoDetails>]}>
                                                <p>视频任务{index}</p>
                                                <p>日期：{item.startDate}-{item.endDate}</p>
                                                <p>学生：{item.avatar}</p>
                                                <p>学科: {item.subject}</p>
                                                <p>描述: {item.content}</p>
                                                <p style={{color: 'red',textAlign:'center'}}>{item.accepted? "已预约":" "}</p>
                                            </Card>
                                        </div>
                                    )
                            })
                        }
                    </div>
                </InfiniteScroll>
            </div>
        )
    }
}
