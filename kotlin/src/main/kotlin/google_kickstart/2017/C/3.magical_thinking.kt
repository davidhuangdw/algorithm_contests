package google_kickstart.`2017`.C

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, Q) = input_ints()
        val A = (1..N).map{split_input()[0]}
        val B = split_input()[0]
        val S = input_ints()

        fun smallSet(): Int {
            var same = 0
            for((a, b) in A[0].toList().zip(B.toList())){
                if(a == b) same ++
            }
            return S[0] + (Q-same) - 2 * maxOf(0, S[0]-same)
        }
        println("Case #${it}: ${smallSet()}")
    }
}
