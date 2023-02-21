import React from "react";
import { useContext } from "react";
import RepoItem from './RepoItem';
import RepoContext from '../../context/repos/RepoContext'
import Spinner from "../layout/Spinner";

function RepoResults() {
    const { repos, loading } = useContext(RepoContext)

    if (!loading) {
        return (
            <div className='grid grid-cols-1 gap-8 xl:grid-cols-4 lg:grid-cols-3 md:grid-cols-2'>
				{repos.map((repo) => (
					<RepoItem key={repo.id} repo={repo}/>
				))}
			</div>
        )
    } else {
        return <Spinner />
    }
}

export default RepoResults
