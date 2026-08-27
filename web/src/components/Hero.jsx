import reactLogo from '../assets/react.svg'
import viteLogo from '../assets/vite.svg'
import heroImg from '../assets/hero.png'

export function Hero() {
  return (
    <div className="hero">
      <img src={reactLogo} className="framework" alt="React logo" />
      <img src={heroImg} className="base" width="170" height="179" alt="" />
      <img src={viteLogo} className="vite" alt="Vite logo" />
    </div>

  )
}

export function Doc() {
  return (
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
  )
}


