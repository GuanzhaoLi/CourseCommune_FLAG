import { Form, Input, InputNumber, Button } from 'antd';
import CheckBox from "./CheckBox";
const layout = {
    labelCol: {
        span: 8,
    },
    wrapperCol: {
        span: 16,
    },
};
const validateMessages = {
    required: '${label} is required!',
    types: {
        email: '${label} is not a valid email!',
        number: '${label} is not a valid number!',
    },
    number: {
        range: '${label} must be between ${min} and ${max}',
    },
};

const PostForm = () => {
    const onFinish = (values) => {
        console.log(values);
    };

    return (
        <Form className="PostForm" {...layout} name="nest-messages" onFinish={onFinish} validateMessages={validateMessages}>
            <Form.Item
                name={['user', 'subject']}
                label="Subject"
                rules={[
                    {
                        required: true,
                    },
                ]}
            >
                <Input />
            </Form.Item>
            <Form.Item
                name={['user', 'age']}
                label="Level"
                rules={[
                    {
                        type: 'number',
                        min: 0,
                        max: 99,

                    },
                ]}
            >
                <CheckBox />
            </Form.Item>

            <Form.Item name={['user', 'question']} label="Question">
                <Input.TextArea />
            </Form.Item>
            <Form.Item wrapperCol={{ ...layout.wrapperCol, offset: 8 }}>
                <Button type="primary" htmlType="submit">
                    Submit
                </Button>
            </Form.Item>
        </Form>
    );
};
export default PostForm;

