package google_kickstart.`2022`.H
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, )= input_ints()
        val A = input_ints()
        val adj = List(N){ mutableListOf<Int>()}
        for(i in 0 until N-1){
            val (a, b) = input_ints()
            adj[a-1].add(b-1)
            adj[b-1].add(a-1)
        }
        val ca = MutableList(N){0}
        fun dp(i: Int): Int{
            if(ca[i] == 0){
                ca[i] = 1
                for(j in adj[i]){
                    if(A[j] < A[i]){
                        ca[i] += dp(j)
                    }
                }
            }
            return ca[i]
        }
        val ma = (0 until N).maxOf { dp(it) }
        println("Case #${it}: $ma")
    }
}
