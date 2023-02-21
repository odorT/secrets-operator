import React from "react";
import { Link } from "react-router-dom";
import PropTypes from 'prop-types'

function RepoItem({ repo: { id, name } }) {
    return (
        <div className='card shadow-md compact side bg-base-100'>
			<div className='flex-row items-center space-x-4 card-body'>
				<div>
					<h2 className='card-title'>{name}</h2>
                    <Link to={`/repos/${id}`} className='text-base-content text-opacity-40'>Check More</Link>
				</div>
			</div>
		</div>
    )
}

RepoItem.propTypes = {
    repo: PropTypes.object.isRequired,
}

export default RepoItem
