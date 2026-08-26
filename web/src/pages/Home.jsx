import { useState, Activity } from 'react'
import reactLogo from '../assets/react.svg'
import viteLogo from '../assets/vite.svg'
import heroImg from '../assets/hero.png'

import '../App.css'
import { Link } from 'react-router-dom'

function Home() {
  const [count, setCount] = useState(0)
  const [hide, setHide] = useState(false)
  return (
    <>
      <section id="center">
        <div className="hero">
          <img src={reactLogo} className="framework" alt="React logo" />
          <img src={heroImg} className="base" width="170" height="179" alt="" />
          <img src={viteLogo} className="vite" alt="Vite logo" />
        </div>
        <div>
          <h1>Socia<span className='hlt'>lx</span></h1>
          <p>
            test my <code>Api</code>
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
          <Link to={"/signup"}>
            <button type="submit" className='button'>Sign Up</button>
          </Link>
          <Link to={"/signin"}>
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
          <ul>
            <li>
              <a href="https://vite.dev/" target="_blank">
                <img className="logo" src={viteLogo} alt="" />
                Explore Vite
              </a>
            </li>
            <li>
              <a href="https://react.dev/" target="_blank">
                <img className="button-icon" src={reactLogo} alt="" />
                Learn more
              </a>
            </li>
          </ul>
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
