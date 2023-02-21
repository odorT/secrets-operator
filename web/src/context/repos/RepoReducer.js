export const ACTIONS = {
    SET_LOADING: 'SET_LOADING',
    GET_REPOS: 'GET_REPOS',
    CLEAR_REPOS: 'CLEAR_REPOS',
    GET_REPO: 'GET_REPO',
    GET_REPOS_AND_FINDINGS: 'GET_REPOS_AND_FINDINGS',
}

function repoReducer(state, action) {
    switch (action.type) {
        case ACTIONS.SET_LOADING:
            return {
                ...state,
                loading: true,
            }
        case ACTIONS.GET_REPOS_AND_FINDINGS:
            return {
                ...state,
                repo: action.payload.repo,
                findings: action.payload.findings,
                loading: false,
            }
        case ACTIONS.GET_REPOS:
            return {
                ...state,
                repos: action.payload,
                loading: false,
            }
        case ACTIONS.CLEAR_REPOS:
                return {
                    ...state,
                    repos: [],
                }
        default:
            return state
    }
}

export default repoReducer
