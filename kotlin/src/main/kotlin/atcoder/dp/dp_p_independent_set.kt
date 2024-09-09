package atcoder.dp
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val MOD = 1000_000_007
    val (N, ) = input_ints()
    val adj = List(N+1) { mutableListOf<Int>() }
    (1..N-1).forEach{
        val (x , y) = input_ints()
        adj[x].add(y)
        adj[y].add(x)
    }

    //build tree
    val chd = List(N+1) { mutableListOf<Int>() }
    fun build(i: Int, par: Int){
        adj[i].filter{it != par}.forEach { j ->
            build(j, i)
            chd[i].add(j)
        }
    }
    build(1, -1)

    // dp
    val DP = Array(N+1){LongArray(2){-1} }
    fun dp(i: Int, t: Int): Long{
        if(DP[i][t] < 0){
            var cnt = 0L

            var cur = 1L
            for(j in chd[i]){
                cur = cur * dp(j, 0) % MOD
            }
            cnt += cur

            if(t == 0){
                cur = 1
                for(j in chd[i]){
                    cur = cur * dp(j, 1) % MOD
                }
                cnt += cur
            }
            DP[i][t] = cnt % MOD
        }
        return DP[i][t]
    }
    println(dp(1, 0))
}
