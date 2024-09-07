import { createSignal, type Component, createEffect, For } from 'solid-js'
import styles from './App.module.css'

interface Pokemon {
  id: string
  name: string
  picture: string
}

const App: Component = () => {
  const [pokemons, setPokemons] = createSignal<Pokemon[]>([])
  const [page, setPage] = createSignal<number>(0)

  createEffect(() => {
    const intersectionObserver = new IntersectionObserver(async (entries) => {
      if (entries[0].intersectionRatio <= 0) {
        return
      }

      const response = await fetch(`/api/pokemon?page=${page()}`)
      const data = await response.json()
      setPokemons((previous) => [...previous, ...data.results])
      setPage((previous) => previous + 1)
    })

    intersectionObserver.observe(document.querySelector('.loading')!)
  })

  return (
    <div id={styles.App}>
      <ul>
        <For each={pokemons()}>
          {(pokemon) => {
            return (
              <li>
                <img src={pokemon.picture} alt={pokemon.name} />
                <div>
                  <span>{pokemon.id.toString().padStart(4, '0')}</span>
                  <span>{pokemon.name}</span>
                </div>
              </li>
            )
          }}
        </For>
      </ul>

      <div class="loading">加载中...</div>
    </div>
  )
}

export default App
