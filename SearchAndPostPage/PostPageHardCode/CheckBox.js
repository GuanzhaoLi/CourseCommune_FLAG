import { Checkbox, Divider } from 'antd';
import React from 'react';
const CheckboxGroup = Checkbox.Group;

const plainOptions = ['Entry', 'Mid', 'Senior'];
const defaultCheckedList = ['Entry', 'Orange'];

const CheckBox = () => {
    const [checkedList, setCheckedList] = React.useState(defaultCheckedList);
    const [indeterminate, setIndeterminate] = React.useState(true);
    const [checkAll, setCheckAll] = React.useState(false);

    const onChange = list => {
        setCheckedList(list);
        setIndeterminate(!!list.length && list.length < plainOptions.length);
        setCheckAll(list.length === plainOptions.length);
    };

    const onCheckAllChange = e => {
        setCheckedList(e.target.checked ? plainOptions : []);
        setIndeterminate(false);
        setCheckAll(e.target.checked);
    };

    return (
        <>


            <CheckboxGroup options={plainOptions} value={checkedList} onChange={onChange} />
        </>
    );
};
export default CheckBox;