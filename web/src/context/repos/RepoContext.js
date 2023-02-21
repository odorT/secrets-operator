import {createContext, useReducer} from 'react'
import repoReducer from './RepoReducer'

const RepoContext = createContext()

export const RepoProvider = ({ children }) => {
    const initialState = {
        repos: [],
        repo: [],
        findings: [],
        loading: false
    }

    const [state, dispatch] = useReducer(repoReducer, initialState)

    return (
        <RepoContext.Provider
            value={{
                ...state,
                dispatch,
            }}
        >
            {children}
        </RepoContext.Provider>
    )
}

export default RepoContext
