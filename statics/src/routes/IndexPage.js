import React from 'react';
import { connect } from 'dva';
import { Card, Table, Button } from 'antd';
import styles from './IndexPage.css';

const columns = [
  {
    title: '创建时间',
    dataIndex: 'created_at',
    key: 'created_at',
  },
  {
    title: 'IP',
    dataIndex: 'ip',
    key: 'ip',
  },
  {
    title: '信息',
    dataIndex: 'message',
    key: 'message',
  },
];

class IndexPage extends React.Component {
  componentDidMount() {
    this.getHistory();
  }

  getHistory() {
    const { dispatch } = this.props;
    dispatch({
      type: 'ip_history/get',
      payload: {
        limit: 10,
        offset: 0,
      }
    });
  }

  handleTableChange(pagination, filters) {
    const { dispatch } = this.props;
    let listFilter = {
      limit: pagination.pageSize,
      offset: (pagination.current - 1) * pagination.pageSize,
    };

    // 获取 midx 列表
    dispatch({
      type: 'ip_history/get',
      payload: listFilter,
    });
  }

  tableView = () => {
    const { ip_history } = this.props;
    if (!ip_history.list) {
      return <></>
    }

    let source = ip_history.list.map((item, index) => {
      return {
        key: index,
        created_at: item.created_at,
        ip: item.ip,
        message: item.message,
      }
    })

    return < Card title="ip 地址" >
      <Table
        dataSource={source}
        rowKey={record => record.id}
        columns={columns}
        pagination={ip_history.pagination}
        onChange={this.handleTableChange}
      />;
    </Card >
  }

  render() {
    const { ip_history } = this.props;
    console.log(ip_history)

    return (
      <div className={styles.normal}>
        <Button onClick={() => this.getHistory()}>获取</Button>
        {this.tableView()}
      </div>
    )
  }
}

IndexPage.propTypes = {
};

export default connect(({ ip_history }) => ({
  ip_history,
}))(IndexPage);