import request from '../utils/request';

export function query() {
  return request('/api/users');
}

export function queryIpHistory(filter) {
  console.log(filter)
  return request('/api/v1/ip_history/get', { method: 'POST', body: JSON.stringify(filter) });
}
