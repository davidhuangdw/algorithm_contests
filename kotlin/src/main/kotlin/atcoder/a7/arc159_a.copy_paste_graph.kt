package atcoder.a7

fun main(){
    fun inputInts() = readLine()!!.split(" ").map { it.toInt() }
    fun inputLongs() = readLine()!!.split(" ").map { it.toLong() }

    val MAX = 1e8.toInt()
    val (N, K) = inputInts()
    val A = (1..N).map {inputInts().map{ if(it > 0) it else MAX}}
    val dist = List(N){i -> IntArray(N){j -> A[i][j]} }

    for(k in 0 until N)
        for(i in 0 until N)
            for(j in 0 until N){
                dist[i][j] = minOf(dist[i][j], dist[i][k] + dist[k][j])
            }

    val (Q) = inputInts()
    repeat(Q){
        var (s, t) = inputLongs()
        val d = dist[((s-1)%N).toInt()][((t-1)%N).toInt()]
        println(if(d < MAX) d else -1)
    }
}
