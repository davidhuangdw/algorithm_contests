package google_code_jam.`2019`.`1a`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, N, M) = input_ints()
    val P = listOf(16, 9, 5, 7, 11, 13, 17)  // trick: only need related prime
    (1..T).forEach {
        val rem = MutableList(P.size){0}
        for((i, p) in P.withIndex()){
            println((1..18).map{p}.joinToString(" "))
             rem[i] = input_ints().sum() % p
        }
        var res = 0
        for(x in 1..M){
            if(P.zip(rem).all{(p, r) -> x % p == r}){
                res = x
                break
            }
        }
        println(res)
        val (r, ) = input_ints()
        if(r != 1)
            throw RuntimeException("failed")
//        println("Case #${it}: $res")
    }
}
