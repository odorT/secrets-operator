import React, { useState, useContext } from "react";
import RepoContext from '../../context/repos/RepoContext'
import AlertContext from "../../context/alert/AlertContext";
import { searchRepos } from '../../context/repos/RepoActions'
import { ACTIONS } from "../../context/repos/RepoReducer";

function RepoSearch() {
	const [text, setText] = useState('');
	const { repos, dispatch } = useContext(RepoContext);
	const { setAlert } = useContext(AlertContext);

	const handleSubmit = async (e) => {
		e.preventDefault();

		if (text === '') {
			setAlert('Please enter something', 'error');
		} else {
			dispatch({type: ACTIONS.SET_LOADING});
			const repos = await searchRepos(text);
			dispatch({type: ACTIONS.GET_REPOS, payload:  repos});
			setText('');
		}
	}

	return (
		<div className='grid grid-cols-1 xl:grid-cols-2 lg:grid-cols-2 md:grid-cols-2 mb-8 gap-8'>
			<div>
				<form onSubmit={handleSubmit}>
					<div className='form-control'>
						<div className='relative'>
							<input
								type='text'
								className='w-full pr-40 bg-gray-200 input input-lg text-black'
								placeholder='Search...'
								value={text}
								onChange={(e) => {
									setText(e.target.value)
								}}
							/>
							<button
								type='submit'
								className='absolute top-0 right-0 rounded-l-none w-36 btn btn-lg'
							>
								Go
							</button>
						</div>
					</div>
				</form>
			</div>
			{repos.length > 0 && (
				<div>
					<button
						className='btn btn-ghost btn-large'
						onClick={() => {
							dispatch({type: ACTIONS.CLEAR_REPOS})
						}}
					>
						Clear
					</button>
				</div>
			)}
		</div>
	)
}

export default RepoSearch;