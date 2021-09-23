
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, ) = input_ints()
        val S = readLine()!!.trim()
        var l = -N
        val dis = MutableList(N){N}
        for((i, ch) in S.withIndex()){
            if(ch == '1')
                l = i
            dis[i] = minOf(dis[i], i-l)
        }
        var r = N*2
        for(i in N-1 downTo 0){
            if(S[i] == '1')
                r = i
            dis[i] = minOf(dis[i], r-i)
        }
        val sum = dis.fold(0L, Long::plus)
        println("Case #${it}: $sum")
    }
}
