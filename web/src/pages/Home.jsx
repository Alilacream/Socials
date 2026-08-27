import { useState, Activity } from 'react'
import '../App.css'
import { Link } from 'react-router-dom'
import { Hero, Doc } from '../components/Hero'
import { useAuth } from '../utils/AuthProvider'
function Home() {
  const [count, setCount] = useState(0)
  const [hide, setHide] = useState(false)
  const { user } = useAuth()
  console.log(user)
  return (
    <>
      <section id="center">
        <Hero />
        <div>
          <h1>Welcome to Socia<span className='hlt'>lx  {user && user.username}</span></h1>
          <p>
            Test Out <code>API</code>
          </p>
        </div>
        <div className='sign-buttons'>
          {/* wahed l ayba t3lmtha f fun facts about react */}

          <Activity mode={hide ? "hidden" : "visible"}>
            <button
              type="button"
              className="button"
              onClick={() => setCount((count) => count + 1)}
            >
              Count is {count}
            </button>
          </Activity>
          <button
            type="button"
            className='button'
            onClick={() => setHide((prev) => !prev)}
          >Hide :3</button>
        </div>

        <div className='sign-buttons'>
          <Link to={"/register"}>
            <button type="submit" className='button'>Sign Up</button>
          </Link>

          <Link to={"/login"}>
            <button type="submit" className='button'>Sign In</button>
          </Link>
        </div>
      </section>

      <div className="ticks"></div>

      <section id="next-steps">
        <div id="docs">
          <svg className="icon" role="presentation" aria-hidden="true">
            <use href="/icons.svg#documentation-icon"></use>
          </svg>
          <h2>Documentation</h2>
          <p>Your questions, answered</p>
          <Doc />
        </div>
        <div id="social">
          <svg className="icon" role="presentation" aria-hidden="true">
            <use href="/icons.svg#social-icon"></use>
          </svg>
          <h2>Connect with me</h2>
          <p>Join the Vite community or Contact me for collaberative projects :P</p>
          <ul>
            <li>
              <a href="https://github.com/Alilacream" target="_blank">
                <svg
                  className="button-icon"
                  role="presentation"
                  aria-hidden="true"
                >
                  <use href="/icons.svg#github-icon"></use>
                </svg>
                GitHub
              </a>
            </li>
            <li>

            </li>
          </ul>
        </div>
      </section>

      <div className="ticks"></div>
      <section id="spacer"></section>
    </>
  )
}

export default Home
