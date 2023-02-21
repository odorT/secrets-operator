import axios from 'axios'

const SECRETS_OPERATOR_URL = 'http://localhost:8080'

const secretsOperator = axios.create({
    baseURL: SECRETS_OPERATOR_URL
})

export const searchRepos = async (text) => {
    const params = new URLSearchParams({
		query: text,
	})

    const response = await secretsOperator.get(`/api/v1/search/repos?${params}`)
    return response.data.items
}

export const getRepoAndFindings = async (id) => {
	const response = await secretsOperator.get(`/api/v1/findings/${id}`)
	return { repo: response.data, findings: response.data.findings }
}
