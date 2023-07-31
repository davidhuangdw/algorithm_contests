package google_kickstart.`2022`.B

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val IMP = "IMPOSSIBLE"
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (R, C) = input_ints()
        val A = (1..R).map{split_input()[0]}
        fun set2(): String{
            if(A[0][0] == '#') return IMP
            val adj = hashMapOf<Pair<Int, Int>, MutableList<Pair<Int, Int>>>()
            val que = ArrayDeque<Pair<Int, Int>>()
            que.add(0 to 0)
            adj[0 to 0] = mutableListOf()
            while(que.isNotEmpty()){
                val (si, sj) = que.removeFirstOrNull()!!
                for((i, j) in listOf(si -1 to sj, si+1 to sj , si to sj-1, si to sj+1)){
                    if(i in 0 until R && j in 0 until C && A[i][j] == '*' && i to j !in adj){
                        adj[i to j] = mutableListOf(si to sj)
                        adj[si to sj]!!.add(i to j)
                        que.add(i to j)
                    }
                }
            }

            val e = hashMapOf<Pair<Int, Int>, HashSet<Pair<Int, Int>>>()
            fun vertexes(i: Int, j: Int) = listOf(i to j, i to j+1, i+1 to j, i+1 to j+1)
            fun removeEdge(a: Pair<Int, Int>, b: Pair<Int, Int>){
                e[a]!!.remove(b)
                e[b]!!.remove(a)
            }
            fun addEdge(a: Pair<Int, Int>, b: Pair<Int, Int>){
                if(a !in e)
                    e[a] = hashSetOf()
                if(b !in e)
                    e[b] = hashSetOf()
                e[a]!!.add(b)
                e[b]!!.add(a)
            }
            var cnt = 0
            for(i in 0 until R)
                for(j in 0 until C){
                    if(A[i][j] == '*'){
                        cnt ++
                        val (a, b, c, d) = vertexes(i*2, j*2)
                        addEdge(a, c)
                        addEdge(a, b)
                        addEdge(b, d)
                        addEdge(c, d)
                    }
                }
            if(adj.size != cnt){
                return IMP
            }


            for((x, l) in adj){
                val (xi, xj) = x
                for(y in l){
                    val (yi, yj) = y
                    if(xi < yi){
                        val (a, b, c, d) = vertexes(xi*2+1, xj*2)
                        addEdge(a, c)
                        addEdge(b, d)
                        removeEdge(a, b)
                        removeEdge(c, d)
                    }else if(xj < yj){
                        val (a, b, c, d) = vertexes(xi*2, xj*2+1)
                        addEdge(a, b)
                        addEdge(c, d)
                        removeEdge(a, c)
                        removeEdge(b, d)
                    }
                }
            }

            fun dir(u: Pair<Int, Int>, v: Pair<Int, Int>): Char{
                val (ui, uj) = u
                val (vi, vj) = v
                return if(ui != vi){
                    if(ui < vi) 'S' else 'N'
                }else{
                    if(uj < vj) 'E' else 'W'
                }
            }

            return buildString{
                val done = hashSetOf(0 to 0)
                var u = 0 to 0
                var succ = true
                while(succ){
                    succ = false
                    for(v in e[u]!!)
                        if(v !in done){
                            done.add(v)
                            append(dir(u, v))
                            u = v
//                        println(u)
                            succ = true
                            break
                        }
                }
                append(dir(u, 0 to 0))
            }
        }
        val res = set2()
        println("Case #${it}: $res")
    }
}
