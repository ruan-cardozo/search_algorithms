# Os algoritmos retornaram os mesmos resultados?
Sim, todos os dois algoritmos (Busca Ingênua, Rabin-Karp) encontraram exatamente 346 ocorrências da palavra "Capitú" no texto. Isso demonstra que, apesar das diferentes abordagens, todos alcançaram o mesmo resultado correto.

# Qual algoritmo foi mais rápido?
O algoritmo mais rápido foi a Busca Ingênua, executando em apenas 817.72µs(aproximadamente 0,81 ms) e o Rabin-Karp foi significativamente mais lento, levando 1.730864ms (cerca de 1,73 ms).
É interessante notar que o algoritmo teoricamente mais simples (Busca Ingênua) superou os algoritmo mais sofisticado neste caso específico.

# O tamanho do texto ou do trecho buscado influencia no tempo de execução?
Sim, ambos têm influência significativa:
- Tamanho do texto: Quanto maior o texto, mais tempo todos os algoritmos levam, pois precisam percorrer mais conteúdo.
- Tamanho do padrão buscado:
    - Na Busca Ingênua, padrões maiores exigem mais comparações em cada posição.
    - No Rabin-Karp, padrões maiores afetam o cálculo do hash.

No exemplo dado, "Capitú" é uma palavra relativamente curta (6 caracteres), o que pode explicar por que a Busca Ingênua teve bom desempenho.

# Em que situações um algoritmo pode ser preferido ao outro?

- Busca Ingênua:
    - Melhor para padrões curtos
    - Ideal para implementações simples
    - Eficiente quando o texto não é muito grande
    - Quando a simplicidade de implementação é mais importante que a performance
- Rabin-Karp:
    - Excelente para buscar múltiplos padrões simultaneamente
    - Útil em detecção de plágio e análise de similaridade
    - Vantajoso quando você precisa encontrar qualquer um dos vários padrões possíveis