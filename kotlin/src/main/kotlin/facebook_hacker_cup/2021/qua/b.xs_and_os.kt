
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val INF = Int.MAX_VALUE
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, ) = input_ints()
        val A = (1..N).map{readLine()!!}
        val blanks = List(2){MutableList(N){-1 to -1} }
        val count = List(2){MutableList(N){0} }
        val dirty = List(2){MutableList(N){false} }
        for(i in 0 until N)
            for(j in 0 until N){
                val ch = A[i][j]
                if(ch == 'X') {
                    count[0][i] += 1
                    count[1][j] += 1
                }else if(ch == 'O'){
                    dirty[0][i] = true
                    dirty[1][j] = true
                }else{
                    blanks[0][i] = i to j
                    blanks[1][j] = i to j
                }
            }

        var min_cost = INF
        var cnt = 0
        val singles = hashSetOf<Pair<Int, Int>>()
        for(t in 0..1)
            for(i in 0 until N)
                if(!dirty[t][i]){
                    val cost = N - count[t][i]
                    if(cost == 1)
                        singles.add(blanks[t][i])
                    if(cost == min_cost)
                        cnt += 1
                    else if(cost < min_cost){
                        min_cost = cost
                        cnt = 1
                    }
                }
        val res = if(min_cost == INF) "Impossible" else if(min_cost == 1) "$min_cost ${singles.size}" else "$min_cost $cnt"
        println("Case #${it}: $res")
    }
}
