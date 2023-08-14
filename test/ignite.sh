ignite scaffold type batch batch_number:uint merkle_root_hash prev_merkle_root_hash zk_proof
ignite scaffold type exelayer validator:array.string voting_power:array.uint latest_batch:uint latest_merkle_root_hash verification_key chain_info id creator

ignite scaffold message add_execution_layer verification_key chain_info --response id
ignite scaffold query show_execution_layer id --response exelayer:Exelayer
ignite scaffold query list_execution_layers --response exelayer:Exelayer --paginated

ignite scaffold message add_batch id batch_number:uint merkle_root_hash prev_merkle_root_hash zk_proof --response batch_status:bool
ignite scaffold query show_batch id batch_number:uint --response batch:Batch

ignite scaffold type vkey id verification_key
ignite scaffold type exelayer_chains creator id:array.string
ignite scaffold query chain_list creator_address --response exelayer_chains:array.string
ignite scaffold query chain_list_detailed creator_address --response chain:Exelayer --paginated
ignite scaffold query verification_key id --response vkey

ignite scaffold query verify id batch_number:uint inputs --response result:bool,message

# TODO: not run yet
ignite scaffold message add_validator_request validator_address --response voting_poll_id

ignite scaffold query list_validators_request --response unvoted_active_voting_poll_ids:array.string
ignite scaffold query list_validators_request --response unvoted_active_voting_poll_details:array.string
ignite scaffold query list_validators_request --response active_voting_poll_ids:array.string
ignite scaffold query list_validators_request --response active_voting_poll_details:array.string

ignite scaffold message add_validator_vote  validator_address

ignite scaffold query list_validators_request

ignite scaffold message add_execution_layer verification_key chain_info --response id