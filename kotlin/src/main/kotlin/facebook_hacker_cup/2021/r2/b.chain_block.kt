
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, ) = input_ints()
        val adj = List(N){mutableListOf<Int>()}
        repeat(N-1){
            val (a, b) = input_ints()
            adj[a-1].add(b-1)
            adj[b-1].add(a-1)
        }
        val F = input_ints()
        val H = MutableList(N){0}
        val P = MutableList(N){-1}

        val que = ArrayDeque<Int>()
        que.add(0)
        while(que.isNotEmpty()){
            val x = que.removeFirst()
            for(y in adj[x])
                if(y != P[x]){
                    P[y] = x
                    H[y] = H[x] + 1
                    que.add(y)
                }
        }

        val fa = MutableList(N){it}
        fun get(x: Int): Int {
            if(fa[x] != x)
                fa[x] = get(fa[x])
            return fa[x]
        }

        val nodes = hashMapOf<Int, MutableList<Int>>()
        for((i, f) in F.withIndex()){
            if(f !in nodes)
                nodes[f] = mutableListOf()
            nodes[f]!!.add(i)
        }
        for((_, arr) in nodes){
            var vs = arr.map{get(it)}.toHashSet().toList()
            if(vs.size == 1) continue
            var first = true
            while(vs.size > 1){
                var nxt = hashSetOf<Int>()
                var min_h = H[vs[0]]
                var max_h = H[vs[0]]
                for(v in vs){
                    min_h = minOf(min_h, H[v])
                    max_h = maxOf(max_h, H[v])
                }
                if(!first) {
                    min_h = minOf(min_h, max_h-1)
                }
                for(v in vs){
                    var u = v
                    while(H[u] > min_h){
                        val p = get(P[u])
                        fa[u] = p
                        u = p
                    }
                    nxt.add(u)
                }
                vs = nxt.toList()
                first = false
            }
        }
        val roots = (0 until N).map{get(it)}.toHashSet()
        println("Case #${it}: ${roots.size - 1}")
    }
}
