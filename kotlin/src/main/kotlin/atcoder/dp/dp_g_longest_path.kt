package atcoder.dp

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (N, M) = input_ints()
    val edge = MutableList(N+1){mutableListOf<Int>()}
    (1..M).forEach{
        val (x, y) = input_ints()
        edge[x].add(y)
    }
    val DP = MutableList(N+1){-1}
    fun dp(x: Int): Int{
        if(DP[x] < 0){
            DP[x] = edge[x].maxOfOrNull { y -> dp(y)+1 } ?: 0;
        }
        return DP[x]
    }
    println((1..N).maxOf{dp(it)})
}
