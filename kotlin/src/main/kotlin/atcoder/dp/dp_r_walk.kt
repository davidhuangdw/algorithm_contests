package atcoder.dp
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val MOD = 1e9.toInt() + 7
    val (NN, K) = input_longs()
    val N = NN.toInt()
    val A = List(N){input_longs()}

    fun matMul(a: List<List<Long>>, b: List<List<Long>>): List<List<Long>> {
        return List(N){i ->
            List(N){j ->
                (0 until N).sumOf { k ->  a[i][k] * b[k][j] % MOD } % MOD
            }
        }
    }

    fun matPow(): List<List<Long>>{
        var a = A
        var res = List(N){i -> List(N){j -> if(i==j) 1L else 0L}}
        var k = K
        while(k > 0){
            if(k % 2 == 1L){
                res = matMul(res, a)
            }
            a = matMul(a, a)
            k /= 2
        }
        return res
    }

    val all = matPow()
//    all.forEach{println(it)}
    val sum = (0 until N).sumOf { i-> (0 until N).sumOf { j-> all[i][j] } % MOD } % MOD
    println(sum)

//    val L = (0 until 10_000).find{ (1L shl it) > K}!!;
//    val G = List(N){List(N){ MutableList(L){0L} } }
//    (0 until N).forEach{i ->
//        input_ints().forEachIndexed{ j, a ->
//            if(a == 1){
//                G[i][j][0] = 1L
//            }
//        }
//    }
//    for(i in 0 until N)
//        for(j in 0 until N)
//            for(l in 1 until L){
//                G[i][j][l] = (0 until N).sumOf { (G[i][it][l-1] * G[it][j][l-1]) % MOD } % MOD
//            }
//
//    val ls = (0 until L).filter{ (1L shl it) and K > 0};
//    val U = ls.size
//    val F = List(N){MutableList(U){0L}}
//    for(u in 0 until U){
//        for(i in 0 until N){
//            val l = ls[u]
//            if(u == 0){
//                F[i][u] = (0 until N).sumOf{G[i][it][l]} % MOD
//            }else {
//                F[i][u] = (0 until N).sumOf{(G[i][it][l]*F[it][u-1]) % MOD} % MOD
//            }
//        }
//    }
//    println(ls)
//    G.forEach{ println("----G ${it}")}
//    F.forEach{ println("----F: ${it}")}
//    println((0 until N).sumOf { F[it][U-1] } % MOD)
}
