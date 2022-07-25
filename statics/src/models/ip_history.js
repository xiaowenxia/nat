
import { queryIpHistory } from './../services/example';

export default {
    namespace: 'ip_history',
    state: {},
    effects: {
        *get({ payload }, { call, put }) {
            const response = yield call(queryIpHistory, payload);
            console.log(response)
            let pagination = {
                current: payload.offset / payload.limit + 1,
                pageSize: payload.limit,
                total: response.data.total,
            }

            yield put({
                type: 'save',
                payload: {
                    list: response.data.List,
                    pagination,
                },
            });
        }
    },
    reducers: {
        'save'(state, { payload }) {
            return {
                ...state,
                list: payload.list,
                pagination: payload.pagination,
            }
        },
    },
};