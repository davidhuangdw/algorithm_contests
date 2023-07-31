package google_code_jam.`2022`.r2

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        var (N) = input_ints()
        val C = (1..N).map { input_longs() }
        val S = (1..(N + 1)).map { input_longs() }
        fun dis(a: List<Long>, b: List<Long>) = (a[0] - b[0]) * (a[0] - b[0]) + (a[1] - b[1]) * (a[1] - b[1])
        fun set1() {
            N = minOf(N, 15)
            val cache = hashMapOf<Pair<Int, Int>, Boolean>()
            val choose = hashMapOf<Pair<Int, Int>, Pair<Int, Int>>()
            fun dp(rem: Pair<Int, Int>): Boolean {
                val (ch, sw) = rem
                if (ch == 0)
                    return true
                if (rem !in cache) {
                    var succ = false
                    for (i in 0 until N) {
                        if ((1 shl i) and ch == 0)
                            continue
                        var min = Long.MAX_VALUE
                        var close = mutableListOf<Int>()
                        for (j in 0 until N + 1) {
                            if ((1 shl j) and sw == 0)
                                continue
                            val d = dis(C[i], S[j])
                            if (d < min) {
                                min = d
                                close = mutableListOf(j)
                            } else if (d == min) {
                                close.add(j)
                            }
                        }
                        for (j in close)
                            if (j != 0 && dp(ch - (1 shl i) to sw - (1 shl j))) {
                                choose[rem] = i to j
                                succ = true
                                break
                            }
                        if (succ)
                            break
                    }
                    cache[rem] = succ
                }
                return cache[rem]!!
            }

            var ch = (1 shl N) - 1
            var sw = (1 shl (N + 1)) - 1
            if (dp(ch to sw)) {
                println("POSSIBLE")
                while (ch > 0) {
                    val (i, j) = choose[ch to sw]!!
                    println("${i + 1} ${j + 1}")
                    ch -= 1 shl i
                    sw -= 1 shl j
                }
            } else {
                println("IMPOSSIBLE")
            }
        }
        fun set2(){
            val D = List(N){i ->
                (0..N).map{dis(C[i], S[it])}
            }
            val adj = List(N+1){ i ->
                val e = mutableListOf<Int>()
                if(i < N){
                    for(j in 1..N)
                        if(D[i][j] <= D[i][0])
                            e.add(j)
                }
                e
            }

            val matchI = MutableList(N){-1}
            val matchJ = MutableList(N+1){N}
            var d = MutableList(N+1){-1}
            fun bfs(): Boolean {
                d = MutableList(N+1){Int.MAX_VALUE}
                val que = ArrayDeque<Int>()
                val visJ = MutableList(N+1){false}
                for(i in 0 until N){
                    if(matchI[i] < 0){
                        que.add(i)
                        d[i] = 0
                    }
                }
                while(que.isNotEmpty()){
                    val i = que.removeFirst()
                    if(d[i] >= d[N])
                        continue
                    for(j in adj[i]){
                        if(visJ[j])
                            continue
                        visJ[j] = true
                        val u = matchJ[j]
                        if(d[i]+1 < d[u]){
                            d[u] = d[i] + 1
                            que.add(u)
                        }
                    }
                }

                return true
            }
        }
        print("Case #${it}: ")
        set1()
    }
}
