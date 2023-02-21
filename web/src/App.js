import React from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import Navbar from './components/layout/Navbar'
import Alert from './components/layout/Alert'
import Home from './components/pages/Home'
import About from './components/pages/About'
import NotFound from './components/pages/NotFound'
import Footer from './components/layout/Footer'
import { RepoProvider } from './context/repos/RepoContext'
import { AlertProvider } from './context/alert/AlertContext'
import Repo from './components/pages/Repo'

function AppRepo() {
	return (
		<RepoProvider>
			<AlertProvider>
				<Router>
					<div className='flex flex-col justify-between h-screen'>
						<Navbar />

						<main className='container mx-auto px-3 pb-12'>
							<Routes>
								<Route
									path='/'
									element={
										<>
											<Alert />
											<Home />
										</>
									}
								/>
								<Route path='/about' element={<About />} />
								<Route path='/repos/:repo' element={<Repo />} />
								<Route path='/notfound' element={<NotFound />} />
								<Route path='/*' element={<NotFound />} />
							</Routes>
						</main>
						<Footer />
					</div>
				</Router>
			</AlertProvider>
		</RepoProvider>
	)
}

export default AppRepo
