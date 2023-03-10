import React from 'react'

function About() {
  return (
    <>
      <h1 className='text-6xl mb-4'>Secrets Operator WEB</h1>
      <p className='mb-4 text-2xl font-light'>
        A React app to search hard coded secrets found by Secrets Operator.
      </p>
      <p className='text-lg text-gray-400'>
        Version: <span> 1.0.0</span>
      </p>
    </>
  )
}

export default About
