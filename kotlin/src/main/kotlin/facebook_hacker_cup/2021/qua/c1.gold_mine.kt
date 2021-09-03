
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, ) = input_ints()
        val C = input_longs()
        val adj = List(N+1){ mutableListOf<Int>()}
        repeat(N-1){
            val (u, v) = input_ints()
            adj[u].add(v)
            adj[v].add(u)
        }
        val vs = hashSetOf(1)
        var (pre, mx) = 0L to 0L
        fun dfs(u: Int, p: Int): Long {
            if(u in vs) return 0L
            vs.add(u)
            var sub = 0L
            for(v in adj[u])
                sub = maxOf(sub, dfs(v, u))
            return sub + C[u-1]
        }
        for(u in adj[1])
            if(u !in vs){
                val cur = dfs(u, 1)
                mx = maxOf(mx, pre+cur)
                pre = maxOf(pre, cur)
            }
        mx += C[0]
        println("Case #${it}: $mx")
    }
}
