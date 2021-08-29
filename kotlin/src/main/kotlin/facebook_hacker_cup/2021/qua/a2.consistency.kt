
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val M = 26
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val S = readLine()!!
        val (K, ) = input_ints()
        val INF = Int.MAX_VALUE
        val dis = List(M){i -> MutableList(M){j -> if(i==j) 0 else INF} }
        repeat(K){
            val (x, y) = readLine()!!.toList()
            dis[x-'A'][y-'A'] = 1
        }
        for(k in 0 until M)
            for(i in 0 until M)
                for(j in 0 until M)
                    if(dis[i][k] < INF && dis[k][j] < INF)
                        dis[i][j] = minOf(dis[i][j], dis[i][k] + dis[k][j])

        var min_cost = INF
        for(y in 0 until M){
            var cnt = 0
            for(ch in S) {
                val x = ch - 'A'
                if (dis[x][y] == INF) {
                    cnt = -1
                    break
                }
                cnt += dis[x][y]
            }
            if(cnt >= 0)
                min_cost = minOf(min_cost, cnt)
        }

        if(min_cost == INF)
            min_cost = -1
        println("Case #${it}: $min_cost")
    }

}
