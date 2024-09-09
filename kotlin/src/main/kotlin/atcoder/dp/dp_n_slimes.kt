package atcoder.dp
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (N, ) = input_ints()
    val A = input_longs()
    val DP = Array(N){LongArray(N)}
    val pre = A.scan(0, Long::plus)

    for(d in 1 until N)
        for(i in 0..N-1-d){
            val j = i+d
            DP[i][j] = pre[j+1] - pre[i] + ((i..j-1).minOfOrNull { DP[i][it]+DP[it+1][j] } ?: 0)
//            println("------- $d: $i $j: ${DP[i][j]}")
        }
    println(DP[0][N-1])
}
