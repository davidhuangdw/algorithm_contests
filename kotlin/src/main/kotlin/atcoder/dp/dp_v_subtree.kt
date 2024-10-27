package atcoder.dp
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (N, M) = input_ints()
    val adj = List(N+1){ mutableListOf<Int>() }
    (0..N-2).forEach {
        val (x, y) = input_ints()
        adj[x].add(y)
        adj[y].add(x)
    }

    var sum = 1L
    fun dfs(v: Int, p: Int): Long {
        var mul = 1L
        for(u in adj[v]){
            if(u != p){
                mul = (mul * (dfs(u, v)+1)) % M
            }
        }
        sum = (sum + mul) % M
        return mul
    }
    println(dfs(1, -1))
//    println(sum)
}
