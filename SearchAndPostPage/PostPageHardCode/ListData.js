import React from 'react';
import {Avatar, List, Space} from "antd";
import {LikeOutlined, MessageOutlined, StarOutlined} from "@ant-design/icons";

function ListData(props) {
    const originData = [{
        QId: 1,
        StartTime: "2021-01-30",
        EndTime: "2021-02-01",
        RequestBy: "John",
        FulfilledBy: "Simon",
        Subject: "Math",
        Level: "EntryLevel",
        KeyWord: "what is the answer of the quadratic equation: x^2 – 2x – 15 = 0",
        Answer: "x = 5 or x = -3",
        S_t_rating: 5
    },
        {
            QId: 2,
            StartTime: "2021-01-29",
            EndTime: "2021-02-01",
            RequestBy: "Mike",
            FulfilledBy: "Simon",
            Subject: "Math",
            Level: "EntryLevel",
            KeyWord: "How to calculate quadratic equation",
            Answer: "The solution(s) to a quadratic equation can be calculated using the Quadratic Formula:\n" +
                "\n" +
                "x = [ -b plus minus square root of (b^2-4ac) ] / 2a",
            S_t_rating: 5
        },
        {
            QId: 3,
            StartTime: "2021-01-29",
            EndTime: "2021-02-02",
            RequestBy: "Jack",
            FulfilledBy: "Simon",
            Subject: "Math",
            Level: "EntryLevel",
            KeyWord: "tips of calculate quadratic equation",
            Answer: "To solve a quadratic equation by factoring,\n" +
                "\n" +
                "Put all terms on one side of the equal sign, leaving zero on the other side.\n" +
                "\n" +
                "Factor.\n" +
                "\n" +
                "Set each factor equal to zero.\n" +
                "\n" +
                "Solve each of these equations.\n" +
                "\n" +
                "Check by inserting your answer in the original equation.\n" +
                "\n" +
                "Example 1\n" +
                "\n" +
                "Solve x 2 – 6 x = 16.\n" +
                "\n" +
                "Following the steps,\n" +
                "\n" +
                "x 2 – 6 x = 16 becomes x 2 – 6 x – 16 = 0\n" +
                "\n" +
                "Factor.\n" +
                "\n" +
                "\n" +
                "( x – 8)( x + 2) = 0\n" +
                "\n" +
                "Setting each factor to zero, equation\n" +
                "\n" +
                "Then to check, equation\n" +
                "\n" +
                "Both values, 8 and –2, are solutions to the original equation.\n" +
                "\n",
            S_t_rating: 5
        }

    ];
    const listData = originData.map(aData => {
        return {
            title: aData.KeyWord,
            description: `RequestBy: ${aData.RequestBy}, FulfilledBy: ${aData.FulfilledBy},  Time: ${aData.StartTime}  to  ${aData.EndTime}`,
            content: aData.Answer,
            rate: aData.S_t_rating
        }
    });
    const IconText = ({icon, text}) => (
        <Space>
            {React.createElement(icon)}
            {text}
        </Space>
    );

    // for (let i = 0; i < 2; i++) {
    //     listData.push({
    //         href: 'https://ant.design',
    //         title: `ant design part ${i}`,
    //         avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
    //         description:
    //             'Ant Design, a design language for background applications, is refined by Ant UED Team.',
    //         content:
    //             'We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.',
    //     });
    // }
    return (
        <List
            itemLayout="vertical"
            size="large"
            pagination={{
                onChange: page => {
                    console.log(page);
                },
                pageSize: 3,
            }}
            dataSource={listData}
            footer={
                <div>

                </div>
            }
            renderItem={item => (
                <List.Item
                    key={item.title}
                    actions={[
                        <IconText icon={StarOutlined} text="5" key="list-vertical-star-o"/>,
                    ]}

                >
                    <List.Item.Meta

                        title={item.title}
                        description={item.description}
                    />
                    {item.content}
                </List.Item>
            )}
        />
    );
}

export default ListData;