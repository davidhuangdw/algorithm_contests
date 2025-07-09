package atcoder.a3

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }

    val MOD = 998244353

    val (N, K) = input_ints()

    var nl = 0
    val ends = (1..K).map {
        val (dir, k) = split_input();
        val d = if (dir == "L") 1 else 0 // 1-L, 0-R
        nl += d
        k.toInt() to d
    }.toMap()

    var nr = K -nl
    nl = 0
    val sum = (1..N).fold(1L){acc, i->
        if(ends.containsKey(i)){
            if(ends[i] == 1){
                nl++
            }else{
                nr--
            }
            acc
        }else{
            acc * (nl + nr) % MOD
        }
    }
    println(sum)
}
