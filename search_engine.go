package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func NaiveSearch(text string, pattern string) []int {
    positions := []int{}
    n := len(text)
    m := len(pattern)
    
    for i := 0; i <= n-m; i++ {
        j := 0

        for j < m && text[i+j] == pattern[j] {
            j++
        }

        if j == m {
            positions = append(positions, i)
        }
    }
    return positions
}

func RabinKarp(text string, pattern string) []int {
    positions := []int{}
    n := len(text)
    m := len(pattern)
    if m > n {
        return positions
    }
    
    // Valores para hashing
    d := 256 // tamanho do alfabeto
    q := 101 // número primo
    
    patternHash := 0 
    textHash := 0
    h := 1

    for i := 0; i < m-1; i++ {
        h = (h * d) % q
    }
    
    for i := 0; i < m; i++ {
        patternHash = (d * patternHash + int(pattern[i])) % q
        textHash = (d * textHash + int(text[i])) % q
    }

    for i := 0; i <= n-m; i++ {

        if patternHash == textHash {
            match := true
            for j := 0; j < m; j++ {
                if text[i+j] != pattern[j] {
                    match = false
                    break
                }
            }
            if match {
                positions = append(positions, i)
            }
        }
        
        if i < n-m {
            textHash = (d*(textHash - int(text[i])*h) + int(text[i+m])) % q

            if textHash < 0 {
                textHash += q
            }
        }
    }
    
    return positions
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {

    if len(os.Args) < 2 {
        fmt.Println("Uso: go run search_engine.go 'padrão a ser buscado'")
        return
    }

    pattern := os.Args[1]

    content, err := ioutil.ReadFile("book.txt")
    if err != nil {
        fmt.Println("Erro ao ler o arquivo:", err)
        return
    }
    
    text := string(content)
    
    startTime := time.Now()
    naive_positions := NaiveSearch(text, pattern)
    naiveTime := time.Since(startTime)
    
    startTime = time.Now()
    rk_positions := RabinKarp(text, pattern)
    rkTime := time.Since(startTime)

    fmt.Println("Resultados da busca para:", pattern)
    
    fmt.Println("\n1. Busca Ingênua (Naive Search):")
    fmt.Printf("   Encontradas %d ocorrências em %v\n", len(naive_positions), naiveTime)
    
    fmt.Println("\n2. Rabin-Karp:")
    fmt.Printf("   Encontradas %d ocorrências em %v\n", len(rk_positions), rkTime)
    
    if len(naive_positions) > 0 {
        fmt.Printf("\nPrimeira ocorrência: posição %d\n", naive_positions[0])
        start := max(0, naive_positions[0]-20)
        end := min(len(text), naive_positions[0]+len(pattern)+20)
        fmt.Printf("Contexto: ...%s...\n", text[start:end])
        
        if len(naive_positions) > 1 {
            fmt.Printf("\nOutras ocorrências nas posições: ")
            for i := 1; i < min(5, len(naive_positions)); i++ {
                fmt.Printf("%d ", naive_positions[i])
            }
            if len(naive_positions) > 5 {
                fmt.Printf("... (total de %d ocorrências)", len(naive_positions))
            }
            fmt.Println()
        }
    } else {
        fmt.Println("\nPadrão não encontrado no texto.")
    }
}