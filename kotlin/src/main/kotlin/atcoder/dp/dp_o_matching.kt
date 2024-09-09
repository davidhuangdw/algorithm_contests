package atcoder.dp
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val MOD = 1000_000_007
    val (N, ) = input_ints()
    val A = Array(N){input_ints().toIntArray()}

    val S = 1 shl N
    val DP = LongArray(S)
    DP[0] = 1
    for(s in 1 until S){
        val i = (0 until N).count{ (1 shl it) and s == 0}
//        println("------ $s: $i")
        var sum = 0L
        for(j in 0 until N){
            if((1 shl j) and s != 0 && A[i][j]==1){
                sum = (sum + DP[s - (1 shl j)])  % MOD
//                println("=====$s: $i: $j: ${s - (1 shl j)} ${DP[s- (1 shl j)]}")
            }
        }
        DP[s] = sum
//        println("------ $s: $i :$sum")
    }
    println(DP[S-1])

}
