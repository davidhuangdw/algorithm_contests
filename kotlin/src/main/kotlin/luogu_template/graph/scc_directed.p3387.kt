package luogu_template.graph

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    val (N, M) = inputInts()
    val edges = List(N) { mutableListOf<Int>() }

    val weight = inputInts()
    (1..M).forEach {
        val (a, b) = inputInts()
        edges[a - 1].add(b - 1)
    }

    var sccCnt = 0
    val sccSum = IntArray(N)
    val scc = IntArray(N)

    val dfn = IntArray(N)
    val low = IntArray(N)
    var time = 0
    val stack = mutableListOf<Int>()
    fun tarjan(u: Int) {
        dfn[u] = ++time
        low[u] = dfn[u]
        stack.add(u)
        for (v in edges[u]) {
            if (dfn[v] < 0) {
                continue
            } // already poped, shouldn't use low[v]
            if (dfn[v] == 0) {
                tarjan(v)
            }
            low[u] = minOf(low[u], low[v])
        }
        if (low[u] == dfn[u]) {
            var sum = 0
            do {
                val v = stack.removeLast()
                scc[v] = sccCnt // bug: scc[v]= u
                dfn[v] = -1 // exclude poped
                sum += weight[v]
            } while (v != u)
            sccSum[sccCnt++] = sum
        }
    }

    // build scc
    for (u in 0 until N)
        if (dfn[u] == 0) {
            tarjan(u)
        }

    // build scc graph
    val sccEdge = List(sccCnt){ hashSetOf<Int>() }
    for(u in 0 until N)
    for(v in edges[u]){
        if(scc[u] != scc[v]){
            sccEdge[scc[u]].add(scc[v])
        }
    }

    // dp
    val dp = IntArray(sccCnt){sccSum[it]} // max-path till scc u
    for(su in sccCnt-1 downTo 0){ // leaf last
        for(sv in sccEdge[su]){
            dp[sv] = maxOf(dp[sv], dp[su]+sccSum[sv])
        }
    }
    println(dp.max())
}
